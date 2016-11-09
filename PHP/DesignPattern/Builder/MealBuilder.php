<?php

/**
 * 抽象建造者类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
abstract class MealBuilder {
    abstract function BuildDrink();

    abstract function BuildFood();

    abstract function getMeal();
}