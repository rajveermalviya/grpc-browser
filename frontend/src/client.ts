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
  const result = document.getElementById("result") as HTMLParagraphElement;

  checkUsernamebtn.addEventListener("click", async () => {
    result.innerHTML = "";
    // create a new request object.
    const checkUsernameRequest = new CheckUsernameRequest();
    // set the username value
    checkUsernameRequest.setUsername(input.value);
    const e = document.createElement("div");

    // and call the gRPC endpoint.
    // with the callback.
    try {
      const res = await client.checkPromise(checkUsernameRequest, {});

      if (res.getIsvalid()) {
        e.innerText = "Username is valid";
      } else {
        e.innerText = "Username is invalid";
      }
      result.append(e);
    } catch (err) {
      console.error(err);
      e.innerText = "Error Occured : " + JSON.stringify(err);
      result.append(e);
    }
  });

  getUserbtn.addEventListener("click", async () => {
    result.innerHTML = "";

    // if (username) e.removeChild(username);
    // if (name) e.removeChild(name);
    // if (image) e.removeChild(image);
    // if (about) e.removeChild(about);
    // e.remove();

    const getUserDetailsRequest = new GetUserDetailsRequest();
    getUserDetailsRequest.setUsername(input.value);

    const e = document.createElement("div");

    try {
      const res = await client.getUserPromise(getUserDetailsRequest, {});
      // pretty print the response in json.
      const username = document.createElement("p");
      const name = document.createElement("p");
      const image = document.createElement("img");
      e.className = "userData";

      username.id = "username";
      name.id = "name";
      image.id = "image";

      username.innerText = "@" + res.getUsername();
      name.innerText = res.getName();
      image.src = res.getImageurl();

      e.append(username, name, image);

      result.append(e);
    } catch (err) {
      console.error(err);
      e.innerText = "Error Occured : " + JSON.stringify(err);
      result.append(e);
    }
  });
});
