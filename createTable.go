package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	//　DynamoDBクライアントの新しいインスタンスを作成します。
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// CreateTableInputはCreateTable操作の入力を表します。
	params := &dynamodb.CreateTableInput{
		// テーブルやインデックスのキースキーマを記述する属性を表します。
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"), // プライマリキー名
				AttributeType: aws.String("N"),  // データ型(String:S, Number:N, Binary:B の三種)
			},
		},
		// キースキーマの単一の要素を表します。 キー・スキーマは、表の主キーを構成する属性、またはインデックスのキー属性を指定します。
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),   // インデックス名
				KeyType:       aws.String("HASH"), // インデックスの型(HASH または RANGE)
			},
		},
		// 指定したテーブルまたはインデックスのプロビジョニングされたスループット設定を表します。
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1), // 読み込みキャパシティーユニット（デフォルト：１）
			WriteCapacityUnits: aws.Int64(1), // 書き込みキャパシティーユニット（デフォルト：１）
		},
		TableName: aws.String("count"), // テーブル名
	}

	// CreateTableは、新しいテーブルをアカウントに追加します。
	resp, err := ddb.CreateTable(params)

	if err != nil {
		fmt.Println(err.Error()) // エラー処理
	}

	fmt.Println(resp)
}
