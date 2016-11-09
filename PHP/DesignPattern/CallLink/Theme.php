<?php

/**
 * 场景主题类(支持链式调用)
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/14
 */
class Theme {
    private $_id;
    private $_title;
    private $_subtitle;
    private $_icon;
    private $_type;

    public function __construct() {
    }

    public function getId() {
        return $this->_id;
    }

    public function setId($id) {
        $this->_id = $id;
        return $this;
    }

    public function getTitle() {
        return $this->_title;
    }

    public function setTitle($title) {
        $this->_title = $title;
        return $this;
    }

    public function getSubtitle() {
        return $this->_subtitle;
    }

    public function setSubtitle($subtitle) {
        $this->_subtitle = $subtitle;
        return $this;
    }

    public function getIcon() {
        return $this->_icon;
    }

    public function setIcon($icon) {
        if (!empty($icon)) {
            $this->_icon = $icon;
        }
        return $this;
    }

    public function getType() {
        return $this->_type;
    }

    public function setType($type) {
        $this->_type = $type;
        return $this;
    }

    public function toArray() {
        return array(
            'id' => $this->_id,
            'title' => $this->_title,
            'subtitle' => $this->_subtitle,
            'icon' => $this->_icon,
            'type' => $this->_type,
        );
    }

    function __toString() {
        return json_encode(array(
            'id' => $this->_id,
            'title' => $this->_title,
            'subtitle' => $this->_subtitle,
            'icon' => $this->_icon,
            'type' => $this->_type,
        ));
    }
}