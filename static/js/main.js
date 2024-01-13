let imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png",
]

function choose( value ) {
    fetch("/play?c=" + value)
    .then( response => response.json() )
    .then( data => {

        if ( value == 0) {
            document.getElementById("player_choice").innerHTML = "El jugador eligó PIEDRA"
        } else if ( value ) {
            document.getElementById("player_choice").innerHTML = "El jugador eligó PAPEL"        
        } else {
            document.getElementById("player_choice").innerHTML = "El jugador eligó TIJERA"        
        }

        // información mostrada
        document.getElementById("player_score").innerHTML = data.player_score        
        document.getElementById("computer_score").innerHTML = data.computer_score

        document.getElementById("computer_choice").innerHTML = data.computer_choice        
        document.getElementById("round_result").innerHTML = data.round_result        
        document.getElementById("round_message").innerHTML = data.message  
        
        // Mostrando imagen de acuerdo a la opción elegida
        var imgElement = document.getElementById("img_computer")
        imgElement.src = imgArray[ data.computer_choice_int ]
    })
}