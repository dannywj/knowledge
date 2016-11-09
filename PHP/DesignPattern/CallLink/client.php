<?php
/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
require_once(__DIR__ . '/Theme.php');

$theme = new Theme();
// 链式调用设置多个属性并返回最终结果
$result = $theme->setId(123)->setIcon('icon test')->setSubtitle('test_subtitle')->setTitle('test')->setType('api')->toArray();
var_dump($result);

// 重写了__toString方法 使之可以直接输出对象
echo $theme;