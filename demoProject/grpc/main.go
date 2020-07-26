package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"zonst/logging"
	pb "zonst/qipai-oversea/apisrv/propscommonapisrv/protobuf/propscommon"
)

func main() {
	r := gin.Default()
	// grpc测试
	r.POST("/1", func(c *gin.Context) {
		//grpc
		//连接服务器
		conn, err := grpc.Dial(":9601", grpc.WithInsecure())
		if err != nil {
			logging.Errorf("grpc连接err:%v", err)
		}
		defer conn.Close()
		client := pb.NewReceiveAwardServerClient(conn)
		var req *pb.BaseAddPropRequest
		if err := c.Bind(&req); err != nil {
			logging.Errorf("grpc参数err:%v", err)
			return
		}
		resp, err := client.AddCoins(context.Background(), req)
		if err != nil {
			logging.Errorf("return err:%v", err)
			return
		}
		logging.Infof("return success:%v", resp)
	})

	// redis设置过期时间
	r.POST("/2", func(c *gin.Context) {
		//获取连接
		conn, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("connect redis error:", err)
			return
		}
		defer conn.Close()
		//1、判断是否存在
		isExistKey, err := redis.Bool(conn.Do("EXISTS", "name"))
		if err != nil {
			fmt.Println("redis exists error:", err)
			return
		}
		if isExistKey {
			fmt.Printf("%v redis exists !!!\n", isExistKey)
			return
		}
		//2、不存在写入
		_, err = conn.Do("SET", "name", true)
		if err != nil {
			fmt.Println("redis set error:", err)
			return
		}
		//3、读取
		name, err := redis.String(conn.Do("GET", "name"))
		fmt.Printf("redis name:%v\n", name)
		//4、设置十五秒后过期
		_, err = conn.Do("EXPIRE", "name", 15)
		time.Sleep(time.Second * 5)
		//5、在五秒后读取是否存在
		isExistKey, err = redis.Bool(conn.Do("EXISTS", "name"))
		if err != nil {
			fmt.Println("redis exists error:", err)
			return
		}
		if isExistKey {
			name, err = redis.String(conn.Do("GET", "name"))
			fmt.Printf("5秒后 redis name:%v\n", name)
		}
		fmt.Printf("5秒后 redis name:%v\n", isExistKey)
		c.JSON(http.StatusOK, "success")
	})
	//计算距离凌晨的时间还有多少秒
	r.POST("/3", func(c *gin.Context) {
		year := time.Now().Format("2006")
		month := time.Now().Format("01")
		day := time.Now().Day() + 1
		tm, _ := time.Parse("01/02/2006", month+"/"+strconv.Itoa(day)+"/"+year)
		dayLongTime := tm.Unix() - (3600 * 8) - time.Now().Unix()
		c.JSON(http.StatusOK, dayLongTime)
	})
	//获取第三方数据
	r.POST("/4", func(c *gin.Context) {
		type RespInfo struct {
			ErrNo  int      `json:"errno" form:"errno"`
			ErrMsg string      `json:"errmsg" form:"errmsg"`
			Data   interface{} `json:"data" form:"data"`
		}

		resp, err := http.Get("http://129.211.112.56:9506/v1/client/gameconfigapisrv/gameconfig/getConfig?game_id=67&type=3")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()
		//拿到数据[]byte
		data, _ := ioutil.ReadAll(resp.Body)
		//
		fmt.Println(string(data))
		var respData RespInfo
		err = json.Unmarshal(data, &respData)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		// 观看广告后金币配置
		type ADWatchCoinConfig struct {
			ReceiveCount int32 `json:"receive_count"`
			GoldCount    int32 `json:"gold_count"`
		}
		a,_:=json.Marshal(respData.Data)
		b:=make([][]ADWatchCoinConfig,0)
		json.Unmarshal(a,&b)
		fmt.Println("1najkhsdshjdhjhb",len(b))
		fmt.Println("2najkhsdshjdhjhb", len(b[0]))
		fmt.Printf("%T，%v\n", respData.Data,respData.Data)
		c.JSON(http.StatusOK, respData)
	})
	r.Run(":8082") //默认8080端口
}
