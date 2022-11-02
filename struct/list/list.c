

#define Maxsize 100 // 最大空间

/*
 * 顺序表采用顺序存储方式, 即逻辑上相邻的数据在计算机内的存储位置也是相邻的;
 * 顺序存储方式, 元素存储是连续的, 中间不允许有空, 可以快速定位第几个元素, 
 * 但是插入和删除时需要移动大量元素; 根据分配空间方法不同, 顺序表可以分为
 * 静态分配和动态分配两种方法;
 *
 *
 * 采用静态分配方法, 定长数组需要预先分配一段固定大小的连续空间, 但是在运算
 * 的过程中, 如合并、插入等操作, 容易超过预分配的空间长度，出现溢出;
 *
 *
 * 采用动态存储方法, 在运算过程中如果发生溢出, 可以另外开辟一块更大的存储空间, 
 * 用以替换原来的存储空间, 从而达到扩充存储空间的目的
 *
 */

// 顺序表静态分配定义
/* typedef struct { */
/* 	int data[Maxsize]; */
/* 	int length; */
/* }SqList; */


// 顺序表动态分配定义
typedef struct {
	int *elem
	int length
}SqList;

// 初始化顺序表
bool InitList(SqList &L) 
{
	L.elem = new int[Maxsize]; // 为顺序表动态分配Maxsize个空间
	if (! L.elem) return false; // 分配空间失败, TODO:  if (! L.elem)

	L.length = 0;
	return true;
}

// 创建
bool CreateList(SqList &L) // 创建一个顺序表L
{
	int x, i = 0;
	cin>>x;
	while(x != -1) // 输入-1时结束
	{
		if(L.length == Maxsize) 
		{
			count<<"顺序表已满"
			return false;
		}
		cin>>x; // 输入一个元素数据
		L.elem[i++] = x; // cin>>x; // 输入一个元素数据
		L.elem[i++]; // 将数据存入第i个位置,然后i++
	}
	return true;
}

// 取值
bool GetElem(SqList L, int i, int &e) 
{
	if(i<1 || i > L.length) return false; // 索性越界
	e=L.elem[i-1]
	return true;
}

