import { model} from "../../../wailsjs/go/models";
import { h } from "vue";
import AcitonEditor from "./action/index.vue";
export function RandomColor(){
    return `#${Math.floor(Math.random()*16777215).toString(16)}`
}

export function randomTagColor() {
    const color = {
      // color: RandomColor(),
    //   textColor: RandomColor(),
      borderColor: RandomColor(),
    }
    return color
  }

  export function showActionEditor () {
    const m = window.$modal.create({
      title: '创建快捷操作',
      preset: 'card',
      style: {
        width: '400px'
      },
      content: () =>
        h(
          AcitonEditor,
          { type: 'primary' },
          () => '关闭'
        )
    })
  
  }