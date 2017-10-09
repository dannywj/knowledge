# !/usr/bin/python
# -*- coding: UTF-8 -*-
from operator import itemgetter
import time
# 读取web日志文件，按用户分组统计网站浏览时间

# 自定义打印函数
def printList(data):
    for item in data:
        print(item)

#当前数据行的有效时间 = min(当前结束 & 下条开始) - 当前开始
def computeTime(itemList):
    i = 0
    maxLength = len(itemList)
    while i < maxLength:
        currentItem = itemList[i]
        beginTime = currentItem['beginT']
        currentEndTime = currentItem['endT']

        log = '当前结束时间：' + str(currentItem['end'])
        log += ' 当前开始时间：' + str(currentItem['begin'])
        if not (i + 1) == maxLength:  # 存在后面的数据
            nextItem = itemList[i + 1]
            nextBeginTime = nextItem['beginT']
            newEnd = min(currentEndTime, nextBeginTime)
            log += ' 下一条起始时间：' + str(nextItem['begin'])
            useTime = str(newEnd - beginTime)
        else:
            log += ' 下一条起始时间不存在'
            useTime = str(currentEndTime - beginTime)

        print(log)
        print('id:' + str(currentItem['id']) + '  ' +
              currentItem['site'] + ' 使用时间：' + useTime + 's')
        i = i + 1



# 程序开始
f = open(u"D:\\web_log.txt", 'r')
dataList = []
userDataList = {}
for line in f.readlines():
    line = line.strip()
    # print(line)
    item = line.split('\t')
    # print(item)
    # 构造单条数据对象（key-value）
    beginTimeArray = time.strptime(item[3], "%Y-%m-%d %H:%M:%S")
    endTimeArray = time.strptime(item[4], "%Y-%m-%d %H:%M:%S")
    # 转换为时间戳:
    beginTimeStamp = int(time.mktime(beginTimeArray))
    endTimeStamp = int(time.mktime(endTimeArray))

    userId = item[1]
    itemObj = {'id': int(item[0]), 'uid': userId, 'site': item[2], 'begin': item[3],
               'end': item[4], 'beginT': beginTimeStamp, 'endT': endTimeStamp}

    # 判断结果数组中是否有该用户的key，没有则初始化一下
    if userId not in userDataList.keys():
        userDataList[userId] = []

    userDataList[userId].append(itemObj)

print('====单用户数组构造完成====')
print(userDataList)


print('====准备分用户进行数据排序====')
for key, val in userDataList.items():
    # if int(key) == 102:
    print('>>>>>>>>>>当前用户：' + str(key))
    # print(val)
    # 按begin时间排序
    val.sort(key=lambda x: (x["beginT"], x["endT"]))
    print('-----按begintime排序后-----')
    printList(val)
    print('----------')
    computeTime(val)
