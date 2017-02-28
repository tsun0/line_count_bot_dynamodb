package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// セッションを持つDynamoDBクライアントの新しいインスタンスを作成
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// UpdateItem操作の入力を表します。
	param := &dynamodb.UpdateItemInput{
		// Stringは、渡された文字列値へのポインタを返します。
		TableName: aws.String("count"), // テーブル名を指定

		// 属性のデータを表します。
		// 各属性値は、名前と値のペアとして記述されます。
		// 名前はデータ型であり、値はデータそのものです。
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String("1"), // キー名を指定
			},
		},

		// 属性の追加します。
		ExpressionAttributeNames: map[string]*string{
			"#total": aws.String("total"), // 項目名をプレースホルダに入れる
		},
		// 属性のデータを表します。
		// 各属性値は、名前と値のペアとして記述されます。
		// 名前はデータ型であり、値はデータそのものです。
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":total_value": {
				N: aws.String("0"), // 値をプレースホルダに入れる
			},
		},

		UpdateExpression: aws.String("set #total = :total_value"), //プレースホルダを利用して更新の式を書く

		//あとは返してくる情報の種類を指定する
		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	}

	// 既存の項目の属性を編集するか、または新しい項目をテーブルに追加します（存在しない場合）。
	resp, err := ddb.UpdateItem(param) //実行

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
