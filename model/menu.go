package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) CreateMenu(menu Menu) error {
	filter := bson.D{{Key: "id", Value: menu.MenuId}}
	update := bson.D{{Key: "$set", Value: menu}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colMenu.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println(err)

	}
	return nil
}

func (m *Model) DeleteMenu(menu Menu) error {
	// menu_id를 기반으로 삭제할 데이터 필터 생성
	filter := bson.M{"menuid": menu.MenuId}

	// 데이터 삭제
	result, err := m.colMenu.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)

	}
	log.Println(result)
	return nil
}

func (m *Model) GetAllMenu() []Menu {

	var menus []Menu

	//대량 검색 후 조회 이후 반환
	cursor, err := m.colMenu.Find(nil, bson.M{})
	if err != nil {

		log.Println(err)

	}
	defer cursor.Close(nil)

	if err := cursor.All(nil, &menus); err != nil {
		log.Println(err)

	}

	return menus
}

func (m *Model) DetailMenu(menuID int) *Menu {

	var menu Menu
	// 메뉴 조회
	err := m.colMenu.FindOne(context.TODO(), bson.M{"menuid": menuID}).Decode(&menu)
	fmt.Println("menu", menu)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")

	}
	return &menu
}

func (m *Model) RecommendMenu(menu Menu) int64 {

	update := bson.M{}

	if menu.Recommendation == true {
		// 업데이트할 필드 및 값 설정
		update = bson.M{
			"$set": bson.M{
				"recommendation": false,
			},
		}
	} else {
		update = bson.M{
			"$set": bson.M{
				"recommendation": true,
			},
		}
	}
	fmt.Println("update", update)
	// 업데이트 쿼리 실행
	filter := bson.M{"menuid": menu.MenuId} // 업데이트할 문서를 선택하는 필터 조건
	result, err := m.colMenu.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}

	return result.ModifiedCount
}
