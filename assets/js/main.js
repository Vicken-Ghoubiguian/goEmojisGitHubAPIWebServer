$(function() {

    //
    $(".main_emoji_image").tooltip({
        
        html: true,
        tooltipClass: "main_tooltip-styling",
        show: {
            effect: "explode",
            delay: 250
        },
        hide: {
            effect: "explode",
            delay: 250
        }
    });
});