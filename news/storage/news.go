package storage

import pb "github.com/supermicah/Protobufs/news/micah/wiki/news"

type NewsService interface {
	IsExist(title string) bool
	Add(news *pb.News)
	//ToCheck(id int64)
	//ToTracking(id int64)
}

// newsService 新闻服务
type newsService struct {
}

func (n *newsService) IsExist(title string) bool {

	return false
}

func (n *newsService) Add(news *pb.News) {
}
