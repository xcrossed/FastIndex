# FastIndex
a fast index server by prune go implement.

## 目标
实现一个支持搜索召回的index server，代替低效的ElasticSearch

## 相关技术点
* 存储
    leveldb
* web
    gin
* 分词
    jieba
* 召回支持
    bool检索
    倒排索引
    bm25
    tf-idf
    语文相似度
    wand算法