import { model} from "../../../wailsjs/go/models";

export function RandomColor(){
    return `#${Math.floor(Math.random()*16777215).toString(16)}`
}

export function randomTagColor() {
    const color = {
      // color: RandomColor(),
    //   textColor: RandomColor(),
      borderColor: RandomColor(),
    }
    console.log(color);
    return color
  }