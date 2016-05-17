(function () {
  var data = {
    width: "300",
    height: "",
    background: "",
    foreground: "",
    message: ""
  };
  var defaultBackground = "DEDEDE";
  var defaultForeground = "555555";
  var baseURL = window.location.href + "svg/";
  
  // Taken from http://stackoverflow.com/questions/9600295/automatically-change-text-color-to-assure-readability
  var invertColor = function (hexTripletColor) {
    var color = hexTripletColor;
    color = color.substring(1);           // remove #
    color = parseInt(color, 16);          // convert to integer
    color = 0xFFFFFF ^ color;             // invert three bytes
    color = color.toString(16);           // convert to hex
    color = ("000000" + color).slice(-6); // pad with leading zeros
    color = "#" + color;                  // prepend #
    return color;
  }
  
  var connectElement = function(elemId) {
    var element = document.querySelector("#" + elemId);
    if (element.type === "text" || element.type === "number") {
      element.addEventListener("click", function (event) {
        event.target.setSelectionRange(0, event.target.value.length);
      });
    }
    element.addEventListener("input", function (event) {
      var value = event.target.value;
      data[elemId] = value;
      if (element.type === "color") {
        var colorDisplay = document.querySelector("#" + elemId + " + label .pics-input-color");
        if (colorDisplay) {
          colorDisplay.style.background = value;
          colorDisplay.style.color = invertColor(value);
          colorDisplay.innerHTML = value.toUpperCase();
        }
      }
      render();
    });
  };
  
  var init = function () {
    connectElement("width");
    connectElement("height");
    connectElement("background");
    connectElement("foreground");
    connectElement("message");
    
    // Auto update copyright year
    [].slice.call(document.querySelectorAll(".copyright-year")).forEach(function (ele) {
      ele.innerText = (new Date()).getFullYear(); 
    });
    // initial render
    render();
  };
  
  var render = function () {
    var renderEl = document.querySelector("#render");
    var colors = "";
    var urlMessage = "";
    var dimensions = "" + data.width;
    var url;
    
    if (data.height) {
      dimensions += "x" + data.height;
    }
    
    if (data.background) {
        colors = "/" + data.background.replace("#", "");
    }
    
    if (data.foreground) {
      colors += (data.background ? "/" : "/#" + defaultBackground + "/") + data.foreground.replace("#", "");
    }
    
    if (data.message) {
      if (!colors) {
        colors = "/" + defaultBackground + "/" + defaultForeground;
      }
      
      urlMessage = "/" + data.message;
    }
    
    // Build the URL
    url = baseURL + dimensions + colors + urlMessage;
    
    renderEl.innerHTML = "<div class='pics-container'><input onClick='this.setSelectionRange(0, this.value.length);' class='pics-url' value='" + url + "'/></div>" +
    "<div class='pics-preview'><img src='" + url + "' alt='Placeholder " + dimensions +"'/></div>";
  };
  
  init();
})();