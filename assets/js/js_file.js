$(function() {

    //

    //
    $(".main_emoji_image").tooltip({
        
        html: true,
        tooltipClass: "main_emoji_image_tooltip_styling",
        show: {
            effect: "explode",
            delay: 250
        },
        hide: {
            effect: "explode",
            delay: 250
        }
    });

    //
    $("#error_dialog").dialog({modal: true});
    
});