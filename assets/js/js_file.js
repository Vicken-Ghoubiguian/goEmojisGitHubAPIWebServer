// Definition of the jQuery function for this web application...
$(function() {

    // To define the common tooltip for respectively github and docker image (to redirect either to the GitHub repos of this web application, or to its Docker image on Docker Hub)... 
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
        position:{
            at:"right",
            my:"center right",
        },
        show: {
            effect: "explode",
            delay: 250
        },
        hide: {
            effect: "explode",
            delay: 250
        }
    });

    // To define the error dialog which displays when a URL not recognized by this web application is entered...
    $("#error_dialog").dialog({
        
        modal: true,
        closeOnEscape: false,
        resizable: false,
        open: function(event, ui) {
            $(".ui-dialog-titlebar-close", ui.dialog || ui).hide();
        }

    // To define the '#error_dialog' dialog's style...
    }).prev(".ui-dialog-titlebar").css("background","red").css("color", "white").css("font-family", "cursive");
    
});

// Definition of the 'deleteNonExistantEmoji' function to delete an non-existant emoji (when the error HTTP 404 occured in the <img> HTML tag)...
function deleteNonExistantEmoji(currentHTMLElement) {

    // To remove the element from the table (the image and its corresponding row in the table)...
    currentHTMLElement.parentNode.remove();

    // To clean entirely the console...
    console.clear();
}

// Definition of the 'copyEmoji' function to copy the whished and selected emoji...
function copyEmoji(currentHTMLElement) {

    // To get the emoji URL from the alt attribute from the emoji image
    var emojiURL = currentHTMLElement.alt;

    // To split the emoji URL as an array on the '/' character base...
    var first_array_of_splitted_GitHub_emoji_url = emojiURL.split("/");

    // To determine the 'first_array_of_splitted_GitHub_emoji_url' array last element index
    var last_element_index = first_array_of_splitted_GitHub_emoji_url.length - 1;

    // To split the last element of the 'first_array_of_splitted_GitHub_emoji_url' array as an array on the '.' character base...
    var second_array_of_splitted_GitHub_emoji_url = first_array_of_splitted_GitHub_emoji_url[last_element_index].split(".");

    // To get the emoji's unicode (which is the first element of the 'second_array_of_splitted_GitHub_emoji_url' array)...
    var emojisUnicode = second_array_of_splitted_GitHub_emoji_url[0];

    // To copy the emoji in the paperweight with conversion of the unicode to emoji before...
    navigator.clipboard.writeText(String.fromCodePoint(parseInt(emojisUnicode, 16)));

    // To indicate that the emoji was copied successfully...
    alert(String.fromCodePoint(parseInt(emojisUnicode, 16)) + " copied...");

    /*var c = '1f1fa-1f1f8'.split('-').map(i => parseInt(i,16))
    var t = String.fromCodePoint(...c)
    console.log(t)*/
}
