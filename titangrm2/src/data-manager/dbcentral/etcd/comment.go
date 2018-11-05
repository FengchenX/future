package etcd

import (
	"context"
	"fmt"
	"grm-service/common"
	"strings"

	"github.com/coreos/etcd/clientv3"

	"grm-service/dbcentral/etcd"
	"grm-service/log"

	"data-manager/types"
)

// 添加评论
func (e DynamicDB) CreateComment(cmt *types.Comment) error {
	key := fmt.Sprintf("%s/%s/%s", etcd.KYE_GRM_COMMENT, cmt.DataId, cmt.Id)

	cmtMap := make(map[string]string, 20)
	cmtMap["create_time"] = cmt.CreateTime
	if cmt.ToUser != nil && len(cmt.ToUser.Id) > 0 {
		cmtMap["to_user"] = cmt.ToUser.Id
	}
	cmtMap["from_user"] = cmt.FromUser.Id
	cmtMap["content"] = cmt.Content
	for k, v := range cmtMap {
		if _, err := e.Cli.Put(context.Background(), key+"/"+k, v); err != nil {
			log.Error("Failed to create comment :", err)
			return err
		}
	}
	return nil
}

// 查询评论
func (e DynamicDB) QueryComments(dataId string) ([]*types.Comment, error) {
	key := fmt.Sprintf("%s/%s", etcd.KYE_GRM_COMMENT, dataId)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		log.Error("Failed to query comment :", err)
		return nil, err
	}

	nodes := make(map[string]*types.Comment, 100)
	var cmtIds []string
	for _, kv := range resp.Kvs {
		value := string(kv.Value)
		keys := strings.Split(string(kv.Key), etcd.KEY_SPLIT)

		cmtId := keys[len(keys)-2]
		if user, ok := nodes[cmtId]; !ok || user == nil {
			nodes[cmtId] = &types.Comment{
				Id:         cmtId,
				DataId:     dataId,
				CreateTime: "",
				ToUser:     nil,
				FromUser:   nil,
				Content:    "",
			}
			cmtIds = append(cmtIds, cmtId)
		}
		switch keys[len(keys)-1] {
		case "create_time":
			nodes[cmtId].CreateTime = value
		case "from_user":
			nodes[cmtId].FromUser = &common.UserInfo{Id: value}
			if user, err := e.DynamicEtcd.GetUserName(value); err == nil {
				nodes[cmtId].FromUser.Name = user
			}
		case "to_user":
			nodes[cmtId].ToUser = &common.UserInfo{Id: value}
			if user, err := e.DynamicEtcd.GetUserName(value); err == nil {
				nodes[cmtId].ToUser.Name = user
			}
		case "content":
			nodes[cmtId].Content = value
		}
	}

	var cmts []*types.Comment
	// 遍历map是无序的所以使用了个slice
	for _, cmt := range cmtIds {
		cmts = append(cmts, nodes[cmt])
	}
	return cmts, nil
}

// 删除评论
func (e DynamicDB) DelComment(dataId, commentId, user string) error {
	key := fmt.Sprintf("%s/%s/%s", etcd.KYE_GRM_COMMENT, dataId, commentId)

	// 获取评论用户
	//userKey := fmt.Sprintf("%s/from_user", key)
	//resp, err := e.Cli.Get(context.Background(), userKey)
	//if err != nil {
	//	log.Println("Failed to get explorer (%s):%v", key, err)
	//	return err
	//}
	//if len(resp.Kvs) > 0 && user != string(resp.Kvs[0].Value) {
	//	return fmt.Errorf(util.TR("Invalid permission to do this action."))
	//}

	// 移除评论
	_, err := e.Cli.Delete(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Error("Failed to delete comment :", err)
		return err
	}
	return nil
}
