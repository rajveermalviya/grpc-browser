import {
  CheckUsernameRequest,
  GetUserDetailsRequest
} from "./proto/heurist_pb";
import { HeuristGrpcClient } from "./proto/HeuristServiceClientPb";

const client = new HeuristGrpcClient("http://localhost:8081", null, null);

window.addEventListener("load", async () => {
  const checkUsernamebtn = document.getElementById("checkUsername");
  const getUserbtn = document.getElementById("getUser");
  const input = document.getElementById("username") as HTMLInputElement;
  const res = document.getElementById("response") as HTMLParagraphElement;

  checkUsernamebtn.addEventListener("click", () => {
    const checkUsernameRequest = new CheckUsernameRequest();
    checkUsernameRequest.setUsername(input.value);

    client.check(checkUsernameRequest, {}, (err, response) => {
      if (err) {
        throw err;
      }
      res.innerText = String(response.getIsvalid());
    });
  });

  getUserbtn.addEventListener("click", () => {
    const getUserDetailsRequest = new GetUserDetailsRequest();
    getUserDetailsRequest.setUsername(input.value);

    client.getUser(getUserDetailsRequest, {}, (err, response) => {
      if (err) {
        throw err;
      }
      res.innerText = JSON.stringify({
        username: response.getUsername(),
        name: response.getName(),
        about: response.getAbout(),
        imgUrl: response.getImageurl()
      });
    });
  });
});
