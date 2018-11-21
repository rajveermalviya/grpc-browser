import {
  CheckUsernameRequest,
  GetUserDetailsRequest
} from "./proto/heurist_pb";
import { HeuristGrpcClient } from "./proto/HeuristServiceClientPb";

// create a gRPC client.
const protocol = window.location.protocol;
const host = window.location.host;

const hostname = `${protocol}//${host}`;
console.log(hostname);
const client = new HeuristGrpcClient(hostname, null, null);

window.addEventListener("load", () => {
  // get all the dom elements
  const checkUsernamebtn = document.getElementById("checkUsername");
  const getUserbtn = document.getElementById("getUser");
  const input = document.getElementById("username") as HTMLInputElement;
  const res = document.getElementById("response") as HTMLParagraphElement;

  checkUsernamebtn.addEventListener("click", () => {
    // create a new request object.
    const checkUsernameRequest = new CheckUsernameRequest();
    // set the username value
    checkUsernameRequest.setUsername(input.value);

    // and call the gRPC endpoint.
    // with the callback.
    client.check(checkUsernameRequest, {}, (err, response) => {
      if (err) {
        console.error(err);
        res.innerText = JSON.stringify(err, null, 4);
      } else {
        res.innerText = String(response.getIsvalid());
      }
    });
  });

  getUserbtn.addEventListener("click", () => {
    const getUserDetailsRequest = new GetUserDetailsRequest();
    getUserDetailsRequest.setUsername(input.value);

    client.getUser(getUserDetailsRequest, {}, (err, response) => {
      if (err) {
        console.error(err);
        res.innerText = JSON.stringify(err, null, 4);
      } else {
        // pretty print the response in json.
        res.innerText = JSON.stringify(
          {
            username: response.getUsername(),
            name: response.getName(),
            about: response.getAbout(),
            imgUrl: response.getImageurl()
          },
          null,
          4
        );
      }
    });
  });
});
