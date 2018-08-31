{{define "header"}}
 
<style>
      /*
 * Base structure
 */

/* Move down content because we have a fixed navbar that is 50px tall */
body {
    padding-top: 50px;
  }
  
  
  /*
   * Global add-ons
   */
  
  .sub-header {
    padding-bottom: 10px;
    border-bottom: 1px solid #eee;
  }
  
  /*
   * Top navigation
   * Hide default border to remove 1px line.
   */
  .navbar-fixed-top {
    border: 0;
  }
  
  /*
   * Sidebar
   */
  
  /* Hide for mobile, show later */
  .sidebar {
    display: none;
  }
  @media (min-width: 768px) {
    .sidebar {
      position: fixed;
      top: 51px;
      bottom: 0;
      left: 0;
      z-index: 1000;
      display: block;
      padding: 20px;
      overflow-x: hidden;
      overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
      background-color: #f5f5f5;
      border-right: 1px solid #eee;
    }
  }
  
  /* Sidebar navigation */
  .nav-sidebar {
    margin-right: -21px; /* 20px padding + 1px border */
    margin-bottom: 20px;
    margin-left: -20px;
  }
  .nav-sidebar > li > a {
    padding-right: 20px;
    padding-left: 20px;
  }
  .nav-sidebar > .active > a,
  .nav-sidebar > .active > a:hover,
  .nav-sidebar > .active > a:focus {
    color: #fff;
    background-color: #428bca;
  }
  
  
  /*
   * Main content
   */
  
  .main {
    padding: 20px;
  }
  @media (min-width: 768px) {
    .main {
      padding-right: 40px;
      padding-left: 40px;
    }
  }
  .main .page-header {
    margin-top: 0;
  }
  
  
  /*
   * Placeholder dashboard ideas
   */
  
  .placeholders {
    margin-bottom: 30px;
    text-align: center;
  }
  .placeholders h4 {
    margin-bottom: 0;
  }
  .placeholder {
    margin-bottom: 20px;
  }
  .placeholder img {
    display: inline-block;
    border-radius: 50%;
  }




  /*!
 * IE10 viewport hack for Surface/desktop Windows 8 bug
 * Copyright 2014-2015 Twitter, Inc.
 * Licensed under MIT (https://github.com/twbs/bootstrap/blob/master/LICENSE)
 */

/*
 * See the Getting Started docs for more information:
 * http://getbootstrap.com/getting-started/#support-ie10-width
 */
@-ms-viewport     { width: device-width; }
@-o-viewport      { width: device-width; }
@viewport         { width: device-width; }


/*loading begin*/
.loading_shade {
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    display: -webkit-box;
    -webkit-box-pack: center;
    -webkit-box-align: center;
    background: rgba(255, 255, 255, .7);
    z-index: 9999
}

.loading_box {
    padding: 30px;
    box-sizing: border-box;
    -webkit-box-sizing: border-box
}

.loading_box .loading {
    width: 100px;
    height: 100px;
    margin: 0 auto;
    background-color: #ff8814;
    border-radius: 100%;
    -webkit-animation: load_scaleout 1s infinite ease-in-out;
    animation: load_scaleout 1s infinite ease-in-out
}

.loading_box .loading_text {
    text-align: center;
    color: #333;
    font-size: .12rem
}

@-webkit-keyframes load_scaleout {
    0% {
        -webkit-transform: scale(0)
    }

    100% {
        -webkit-transform: scale(1);
        opacity: 0
    }
}

@keyframes load_scaleout {
    0% {
        transform: scale(.1);
        -webkit-transform: scale(.1)
    }

    100% {
        transform: scale(1.5);
        -webkit-transform: scale(1.5);
        opacity: 0
    }
}

.loading_oneline_box {
    font-size: .12rem;
    color: #FFF;
    text-align: center
}

/*loading end*/
</style>
{{end}}