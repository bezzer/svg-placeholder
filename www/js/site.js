(function () {
  var data = {
    width: "300",
    height: "",
    background: "",
    foreground: ""
  };
  
  var baseURL = window.location.href + "svg/";
  
  var connectElement = function(elemId) {
    var element = document.querySelector("#" + elemId);
    element.addEventListener("input", function (event) {
       data[elemId] = event.target.value;
       render();
    });
  };
  
  var init = function () {
    connectElement("width");
    connectElement("height");
    connectElement("background");
    connectElement("foreground");
    // initial render
    render();
  };
  
  var render = function () {
    var renderEl = document.querySelector("#render");
    var colors = "";
    var dimensions = "" + data.width;
    var url;
    
    if (data.height) {
      dimensions += "x" + data.height;
    }
    
    if (data.background) {
        colors = "/" + data.background.replace("#", "");
        if (data.foreground) {
          colors += "/" + data.foreground.replace("#", "");
        }
    }
    
    // Build the URL
    url = baseURL + dimensions + colors;
    
    renderEl.innerHTML = "<div class='pics-container'><div class='pics-url'>" + baseURL + dimensions + colors + "</div></div>" +
    "<div class='pics-preview'><img src='" + url + "' alt='Placeholder " + dimensions +"'/></div>";
  };
  
  init();
})();