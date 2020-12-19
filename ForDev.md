## For Developers
### requirements
 - docker
 - docker-compose
Linux上での実行を推奨します。WSL,WSL2上での動作は確認していません。
### Usage
`git clone https://github.com/FunLocks/team2001.git`
#### Setting API key
Google Geocoding APIとGoogle Map Javascript APIが使用可能なAPI KEYを用意してください。
`echo {API key } >> back/app/apikey/api.txt`
`echo REACT_APP_GOOGLE_API_KEY ={API key } >> front/React/mapviewer/.env`
#### start containers
team2001/で、
`docker-compose build`
`docker-compose run react "cd /app/mapviewer && npm install"`
`docker-compose up`
を起動してください。
[ahchoo viewer](localhost:8081) でアプリが見れることを確認してください。
これであなたも開発に参加することができます。
## API Documantation
 - server : localhost:8080
### /ahchoo/get
- method : GET
 - return : 最新のAHCHOO DATAをJSON形式で返します。
### /ahchoo/getall
 - method : GET
 - return : すべてのAHCHOO DATAをJSON形式で返します。
### /ahchoo/one-hour
 - method : GET
 - return : 一時間のAHCHOO DATAをJSON形式で返します。
### /ahchoo/one-day
 - method : GET
 - return : 一日のAHCHOO DATAをJSON形式で返します。
### /ahchoo/seven-days
 - method : GET
 - return : 一週間のAHCHOO DATAをJSON形式で返します。
### /ahchoo/thiry-days
 - method : GET
 - return : 一ヶ月のAHCHOO DATAをJSON形式で返します。
thirtyではないことに注意してください。@たつおが悪い。
### /ahchoo/post
 - method : POST
 - return : HTTP status code
#### post時のJSON DATA
```
{
    "latitude": value,
    "longitude": value,
    "air-pressure": value, //Enterprise のみ
    "temp" : value,  //Enterprise のみ
}














