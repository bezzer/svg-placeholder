(function () {
  var data = {
    width: "300",
    height: "",
    backgroundStart: "",
    backgroundEnd: "",
    foregroundText: "",
    foregroundBorder: "",
    message: ""
  };
  var defaultBackground = "DEDEDE";
  var defaultForeground = "555555";
  var location = window.location;
  var baseURL = location.protocol + "//" + window.location.hostname + (location.port? ":" + location.port : "") + "/svg/";
  
  // Connect elemensts based on ID
  var connectElement = function(elemId) {
    var element = document.querySelector("#" + elemId);
    if (element.type === "text" || element.type === "number") {
      element.addEventListener("click", function (event) {
        event.target.setSelectionRange(0, event.target.value.length);
      });
    }
    
    var updateValue = function (value) {
      var current = data[elemId];
      if (current !== value) {
        data[elemId] = value;
        render();
      }
    }
   
    // Update on input change
    element.addEventListener("change", function (event) {
      updateValue(event.target.value);
    });
    
    element.addEventListener("input", function (event) {
      updateValue(event.target.value);
    });
    
    var clearButton = document.querySelector("." + elemId + " .clear");
    if (clearButton) {
      clearButton.addEventListener('click', function (e) {
        e.preventDefault();
        e.stopPropagation();
        element.value = "";
        data[elemId] = "";
        var colorPicker = document.querySelector('.' + elemId);
        if (colorPicker) {
          colorPicker.style.backgroundColor = "";
          colorPicker.style.color = "";
        }
        render();
      });
    }
    
    // Initial update
    updateValue(element.value);
  };
  
  var init = function () {
    connectElement("width");
    connectElement("height");
    connectElement("backgroundStart");
    connectElement("backgroundEnd");
    connectElement("foregroundText");
    connectElement("foregroundBorder");
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
    
    if (data.backgroundStart) {
        colors = "/" + data.backgroundStart.replace("#", "");
        if (data.backgroundEnd) {
          colors += "-" + data.backgroundEnd.replace("#", "");
        }
    }
    
    if (data.foregroundText) {
      colors += (data.backgroundStart ? "/" : "/" + defaultBackground + "/") + data.foregroundText.replace("#", "");
      if (data.foregroundBorder) {
        colors += "-" + data.foregroundBorder.replace("#","");
      }
    }
    
    if (data.message) {
      if (!colors) {
        colors = "/" + defaultBackground + "/" + defaultForeground;
      }
      
      urlMessage = "/" + encodeURIComponent(data.message);
    }
    
    // Build the URL
    url = baseURL + dimensions + colors + urlMessage;
    
    renderEl.innerHTML = "<div class='pics-container'><input onClick='this.setSelectionRange(0, this.value.length);' class='pics-url' value='" + url + "'/></div>" +
    "<div class='pics-preview'><img src='" + url + "' alt='Placeholder " + dimensions +"'/></div>";
  };
  
  init();
})();