REACT_APP_BALLOT_ENDPOINT="http://roost-controlplane:32462"
var URL = REACT_APP_BALLOT_ENDPOINT
var resultData = null; 
var data = {
    results: [
       {
           id: 0,
           name: "Roost",
           total_votes: 1000,
       },
       {
           id: 1,
           name: "Docker",
           total_votes: 200,
       },
       {
            id: 2,
            name: "Minikube",
            total_votes: 300,
        },
        {
             id: 3,
             name: "K3d",
             total_votes: 200,
         },
         {
            id: 4,
            name: "kind",
            total_votes: 300,
        }
   ],
   status:{
      code: 200,
      message: "vote saved successfully"
   }
}

function updateResult() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        console.log(this.status)
        console.log("this: ", this)
 
        if(this.status === 0) {
            out = '<h1>Ballot service is not available.</h1>';
            document.getElementById("error_msg").innerHTML = out;
            document.getElementById("main_content").style.display = "none";
            document.getElementById("error_msg").style.display = "block";
        }

        console.log("readyState:", this.readyState);
        if (this.readyState == 4 && this.status == 200) {
            resultData = JSON.parse(this.responseText)
            console.log("got resultData", resultData)
            updateResultX()
            // return resultData;
        }
     };
        xhttp.open("GET", URL, true);
        xhttp.send();

}

function updateResultX() {
    // getResult()
    // if(resultData == null) {
    //     console.log("getting into null check")
    //     // resultData = data;
    // }
    console.log(resultData.status);
    if(resultData.status.code == 204) {
        console.log(resultData.status.message)
        out = '<h1>'+ resultData.status.message +'</h1>';
        document.getElementById("error_msg").innerHTML = out;
        document.getElementById("main_content").style.display = "none";
        document.getElementById("error_msg").style.display = "block";
    }
    else { //if (resultData.status.code == ) {
        console.log("got here")
        resultData.results.sort((a, b) => {               
            return b.total_votes - a.total_votes;
        });
    
        let total = 0;
        let out = '';
        resultData.results.forEach(item => {
            total += item.total_votes
        });
    
        console.log("total_votes: " + total)
    
        resultData.results.forEach(item => {
            percent = Math.round((100 * item.total_votes) / total);
            src = "";
            if(percent < 1) {
                percent = 0
            }
            
            if(item.id == 0) {
                src = "./assets/roost.png"
            }
            else if(item.id == 1) {
                src = "./assets/docker.png"
            }
            else if(item.id == 2) {
                src = "./assets/minikube.png"
            }
            else if(item.id == 3) {
                src = "./assets/k3d.svg"
            }
            else if(item.id == 4) {
                src = "./assets/kind.png"
            }
            console.log("src: ", src)


            out +=  '  <div class="card">'+
                    '  <div class="cardBackgroundContainer">'+
                    '    <div class="cardBackground"></div>'+
                    '    <div class="cardBackgroundImage">'+
                    '        <img'+
                    '          src='+ src +
                    '          width="150px"'+
                    '          height="150px"'+
                    '          class="image"'+
                    '        />'+
                    '    </div>'+
                    '  </div>'+
                    '  <div class="cardContent"> '+ item.name +
                    '      <div class="progressbar_back">'+
                    '          <div class="progressbar_front" style="width:' + percent + '%"></div> '+
                    '          <div>' + percent + '%</div>'+
                    '      </div>'+
                    '  </div>'+
                    '</div>';
            
        });
        document.getElementById("candidates").innerHTML = out;
        document.getElementById("main_content").style.display = "block";
        document.getElementById("error_msg").style.display = "none";

    }
    
}

date = new Date();
document.getElementById("heading_date").innerHTML = "Developer prefer the below tool for building K8S cluster<br>(As of Date:" + date.toLocaleString() + ")";

updateResult()

// 0: roost
// 1: docker
// 2: minikube
// 3: k3d
// 4: kind 