package handlers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

const Lrc1 = `[ti:最后一页]
[ar:江语晨]
[by:PandaMusic]
[00:00.00]最后一页 - 江语晨 (Jessie Chiang)
[00:03.60]词：宋健彰
[00:07.20]曲：詹宇豪
[00:10.80]编曲：痞克四
[00:14.41]雨停滞天空之间
[00:17.70]像泪在眼眶盘旋
[00:21.10]这也许是最后一次见面
[00:27.99]沿途经过的从前
[00:31.46]还来不及再重演
[00:35.64]拥抱早已悄悄冷却
[00:42.15]海潮声 淹没了离别时的黄昏
[00:49.02]只留下不舍的体温
[00:55.83]星空下 拥抱着快凋零的温存
[01:02.24]爱只能在回忆里完整
[01:12.73]想把你抱进身体里面
[01:17.56]不敢让你看见
[01:20.99]嘴角那颗没落下的泪
[01:26.59]如果这是最后的一页
[01:31.23]在你离开之前
[01:34.67]能否让我把故事重写
[02:07.89]雨停滞天空之间
[02:11.09]像泪在眼眶盘旋
[02:14.58]这也许是最后一次见面
[02:21.37]沿途经过的从前
[02:24.77]还来不及再重演
[02:29.03]拥抱早已悄悄冷却
[02:35.53]海潮声 淹没了离别时的黄昏
[02:42.31]只留下不舍的体温
[02:49.15]星空下 拥抱着快凋零的温存
[02:55.62]爱只能在回忆里完整
[03:02.52]想把你抱进身体里面
[03:07.15]不敢让你看见
[03:10.60]嘴角那颗没落下的泪
[03:16.20]如果这是最后的一页
[03:21.00]在你离开之前
[03:24.35]能否让我把故事重写
[03:30.20]想把你抱进身体里面
[03:34.70]不敢让你看见
[03:38.12]嘴角那颗没落下的泪
[03:43.55]如果这是最后的一页
[03:48.86]在你离开之前
[03:52.28]能否让我把故事重写`

const Lrc2 = `[ti:凄美地]
[ar:郭顶]
[al:飞行器的执行周期]
[by:]
[offset:0]
[00:00.00]凄美地 - 郭顶
[00:01.12]词：郭顶
[00:02.24]曲：郭顶
[00:03.36]曾经我是不安河水
[00:06.43]
[00:07.38]穿过森林误入你心
[00:10.08]没计划扎营
[00:12.00]搁下了是非
[00:14.04]一去不回
[00:15.69]
[00:18.76]如今我是造梦的人吶
[00:21.60]
[00:22.31]怅然若失流连忘返啊
[00:25.81]等潮汐来临
[00:27.36]我就能记起
[00:29.44]你的样子
[00:30.81]
[00:32.42]我没看过
[00:37.61]平坦山丘
[00:41.17]怎么触摸
[00:44.93]开花沼泽
[00:48.80]嘿 等我找到你
[00:54.23]试探你眼睛
[00:58.08]心无旁骛地 相拥
[01:03.08]那是我 仅有的温柔
[01:08.20]也是我爱你的原因
[01:12.55]在这凄美地
[01:15.58]
[01:31.75]曾经这里是无人之地
[01:34.84]
[01:35.40]为何没留下有效地址
[01:38.51]肆意的消息
[01:40.34]迷失在十月
[01:42.27]没有音讯
[01:46.88]如今这里是风和日丽
[01:50.10]
[01:50.66]等你再回来雨过迁徙
[01:53.82]看夜幕将近
[01:55.64]我又能记起
[01:57.66]你的样子
[01:59.24]
[02:00.67]我还记得
[02:05.92]平坦山丘
[02:09.41]如今身在
[02:13.25]开花沼泽
[02:17.10]嘿等我找到你
[02:22.61]试探你眼睛
[02:26.33]心无旁骛地 相拥
[02:31.31]那是我 仅有的温柔
[02:36.51]也是我爱你的原因
[02:40.96]在这凄美地
[02:43.94]
[02:45.17]在这之前
[02:46.35]
[02:46.90]别说再见
[02:48.24]
[02:48.97]我已再经不起离别
[02:51.99]
[02:52.59]在这之前
[02:54.02]
[02:54.53]别说再见
[02:55.87]
[02:56.59]我已经开始了想念
[02:59.70]
[03:00.23]在这之前
[03:01.64]
[03:02.25]别说再见
[03:04.14]请帮我停住这时间
[03:07.61]就这样 别安慰
[03:14.74]嘿 等我找到你
[03:20.15]望住你眼睛
[03:24.09]心无旁骛地 相拥
[03:28.89]那是我 仅有的温柔
[03:34.29]也是我爱你的原因
[03:38.57]如此不可及
[03:40.48]
[03:42.45]如此不思议
[03:44.63]
[03:45.24]让我坠落
[03:47.09]在这凄美地`

type Song struct {
	Source      string `json:"source"`
	SingerPhoto string `json:"singerPhoto"`
	Lyric       string `json:"lyric"`
	Singer      string `json:"singer"`
	Name        string `json:"name"`
}

func GetSongs(context *gin.Context) {
	// json := gin.H{"source": "../../../../public/audio/最后一页.aac", "singerPhoto": "../../../../public/audio/江语晨.webp"}
	list := []Song{
		{Source: "../../../../public/audio/最后一页.aac", SingerPhoto: "../../../../public/audio/江语晨.webp", Lyric: Lrc1, Singer: "江语晨", Name: "最后一页"},
		{Lyric: Lrc2, Singer: "郭顶", SingerPhoto: "../../../../public/audio/郭顶.webp", Source: "../../../../public/audio/凄美地.mp3", Name: "凄美地"},
	}
	context.JSON(http.StatusOK, list)
}
