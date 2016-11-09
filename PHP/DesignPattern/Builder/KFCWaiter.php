<?php

/**
 * 指挥者类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
class KFCWaiter {
    // 建造过程固定不变
    public function setMealBuilder(MealBuilder $builder) {
        $builder->BuildDrink();// 过程1
        $builder->BuildFood();// 过程2
    }
}