$(function() {

    //
    $(".footer_image").tooltip({
        html: true,
        tooltipClass: "footer_image",
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
    $(".main_app_presentation").tooltip({

        html: true,
        tooltipClass: "main_app_presentation_tooltip_styling",
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
    $("#error_dialog").dialog({
        
        modal: true,
        closeOnEscape: false,
        open: function(event, ui) {
            $(".ui-dialog-titlebar-close", ui.dialog || ui).hide();
        }
        
    }).prev(".ui-dialog-titlebar").css("background","red").css("color", "white").css("font-family", "cursive");
    
});

//
function deleteNonExistantEmoji(currentHTMLElement) {

    //
    currentHTMLElement.parentNode.remove();

    //
    console.clear();
}