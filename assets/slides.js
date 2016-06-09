"use strict";

var currentSlide = 0;
document.addEventListener("click", function() {console.log("Hi");}, false);

function scrollTo(slideNr) {
  var slides = document.getElementsByClassName("slide");
  var len = slides.length;
  if (slideNr < 0 || slideNr > len -1) {
    return;
  }
  for (var idx=0; idx < len; idx++) {
    var slide = slides[idx];
    if (idx === slideNr) {
      slide.style.visibility = "visible";
      slide.style.opacity = 1;
    } else {
      slide.style.visibility = "hidden";
      slide.style.opacity = 0;
    }
  }
  currentSlide = slideNr;
}

function scroll(skip) {
  scrollTo(currentSlide + skip);
}

window.onkeydown = function(event) {
  switch (event.key) {
    case "ArrowLeft":
      scroll(-1);
      break;
    case "ArrowRight":
      scroll(1);
      break;
    default:
      console.log(event.key)
  }
}

scrollTo(0);
