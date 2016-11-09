<?php
/**
 * 入口
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
require_once(__DIR__ . '/KFCWaiter.php');
require_once(__DIR__ . '/Meal.php');
require_once(__DIR__ . '/MealBuilder.php');
require_once(__DIR__ . '/SubMealBuilderA.php');
require_once(__DIR__ . '/SubMealBuilderB.php');

// 指挥者
$waiter = new KFCWaiter();
// 建造者A
$builderA = new SubMealBuilderA();
$waiter->setMealBuilder($builderA);
$foo1 = $builderA->getMeal();
// 建造者B
$builderB = new SubMealBuilderB();
$waiter->setMealBuilder($builderB);
$foo2 = $builderB->getMeal();

var_dump($foo1, $foo2);