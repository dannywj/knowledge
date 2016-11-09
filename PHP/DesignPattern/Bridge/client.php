<?php
/**
 * 客户端
 * 变继承为组合的一种方式
 * Bridge 模式把两个角色之间的继承关系改为了耦合的关系，从而使这两者可以从容自若的各自独立的变化
 *
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/1
 */
require_once(__DIR__ . '/Car.php');
require_once(__DIR__ . '/Road.php');
require_once(__DIR__ . '/AutoCar.php');
require_once(__DIR__ . '/Bus.php');
require_once(__DIR__ . '/SpeedRoad.php');
require_once(__DIR__ . '/StreetRoad.php');

// 高速公路上行驶小汽车
$road1 = new SpeedRoad();
// 设置路上行驶的车辆（可变化）
$road1->setCar(new AutoCar());
$road1->Go();
echo '===========';

// 街道上行驶公交车
$road2 = new StreetRoad();
$road2->setCar(new Bus());
$road2->Go();


/*
 * result：
E:\DannyCode\study\Bridge\AutoCar.php:11:string 'auto car run' (length=12)
E:\DannyCode\study\Bridge\SpeedRoad.php:14:string 'on SpeedRoad' (length=12)
===========
E:\DannyCode\study\Bridge\Bus.php:11:string 'Bus run' (length=7)
E:\DannyCode\study\Bridge\StreetRoad.php:12:string 'on StreetRoad' (length=13)
*/