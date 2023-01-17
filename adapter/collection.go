package adapter

type Collection interface {
	Item(id string) // 单个数据
	List()          // 列表数据
	Convert()       // 转换器，构造数据
}

func init() {

}
