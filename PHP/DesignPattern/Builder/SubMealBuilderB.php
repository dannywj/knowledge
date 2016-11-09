<?php

/**
 * 具体建造者类B
 * 实现抽象建造者类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
class SubMealBuilderB extends MealBuilder {
    private $meal;

    public function __construct() {
        $this->meal = new Meal();
    }

    function BuildDrink() {
        $this->meal->setDrink('build drink B');
    }

    function BuildFood() {
        $this->meal->setFood('build food B');
    }

    function getMeal() {
        return $this->meal->getDrink() . '+++' . $this->meal->getFood();
    }
}