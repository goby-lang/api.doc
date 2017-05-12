var elementPosition = $('.sidebar').offset();

$(document).ready(function(){
  $(window).scroll(function(){
    if($(window).scrollTop() > elementPosition.top){
      $('.sidebar').css('position','fixed').css('top','0');
    } else {
      $('.sidebar').css('position','static');
    }    
  });
});
