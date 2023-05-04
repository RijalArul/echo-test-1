{
	"info": {
		"_postman_id": "e25b4450-f5ac-405a-8d5c-e490dd190873",
		"name": "Test-Kstyle",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Get All Members",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:9000/members",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"members"
					]
				}
			},
			"response": []
		},
		{
			"name": "Register as Member",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "username",
							"value": "Arul16",
							"type": "text"
						},
						{
							"key": "gender",
							"value": "Male",
							"type": "text"
						},
						{
							"key": "skin_type",
							"value": "Normal Skin",
							"type": "text"
						},
						{
							"key": "skin_color",
							"value": "brown",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:9000/members/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"members",
						"register"
					]
				}
			},
			"response": []
		},
		{
			"name": "Detail Update Member with Params memberID",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "username",
							"value": "Rijal13",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:9000/members/update/6",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"members",
						"update",
						"6"
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete Member with Params memberID",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "http://localhost:9000/members/delete/7",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"members",
						"delete",
						"7"
					]
				}
			},
			"response": []
		},
		{
			"name": "Get All Products",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:9000/products",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"products"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create New Products",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "name_product",
							"value": "Skin Care - D",
							"type": "text"
						},
						{
							"key": "Price",
							"value": "250000",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:9000/products/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"products",
						"register"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create a reviews product with Params ProductID and header member_id ",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "member_id",
						"value": "1",
						"type": "text"
					}
				],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "desc_review",
							"value": "Halo Palembang",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:9000/reviews/create/5",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"reviews",
						"create",
						"5"
					]
				}
			},
			"response": []
		},
		{
			"name": "Get Reviews Products",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:9000/reviews/5",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"reviews",
						"5"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create Reaction of Reviews with Query DISLIKE or LIKE. Headers member_id and Param review_id",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "member_id",
						"value": "3",
						"type": "text"
					}
				],
				"url": {
					"raw": "http://localhost:9000/like_reviews/3?reaction=DISLIKE",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"like_reviews",
						"3"
					],
					"query": [
						{
							"key": "reaction",
							"value": "DISLIKE"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Get Detail Of Product",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:9000/products/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "9000",
					"path": [
						"products",
						"1"
					]
				}
			},
			"response": []
		}
	]
}