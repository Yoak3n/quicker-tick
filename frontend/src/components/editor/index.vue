<template>
  <div class="editor-wrapper">
    <editor-content :editor="rich" />
    <div class="submit-button" @click="submit">
      <button>提交</button>
      <button @click="test">test</button>
    </div>
  </div>
</template>
<script setup lang="ts">
import {  ref,onMounted, onUnmounted,onActivated, PropType } from "vue";
import { Editor, EditorContent } from '@tiptap/vue-3'
import {RichText} from "@/utils/rich_text"
import { SetEditorContent } from ".";
import { Task } from "@/types";
import { ConvertTaskView,ConvertTask } from "../../../wailsjs/go/app/App";
import { model} from "../../../wailsjs/go/models";
let rich =ref<Editor | undefined>()
let result = ref<string>('')

onMounted(()=>{
  rich.value = RichText
  rich.value?.on('update', ({editor }) => {
      const content = JSON.stringify(editor.getJSON())
      result.value = content;
    })
  // const content = ref<Content>([])
  if (props.data != undefined){
    ConvertTaskView(props.data as model.TaskView[]).then((content)=>{
      const a = JSON.stringify(content)
      console.log(a);
      ConvertTask(content).then((doc)=>{
        console.log(doc);
        
        SetEditorContent(rich.value!,doc)
      })
    })
  }  
})

const test = ()=>{
  if (props.data != undefined){
    ConvertTaskView(props.data as model.TaskView[]).then((content)=>{
      const a = JSON.stringify(content)
      console.log(a);
      ConvertTask(content).then((doc)=>{
        console.log(doc);
        SetEditorContent(rich.value!,doc)
      })
    })
  }  
  // const content = `        <ul data-type="taskList">
  //         <li data-type="taskItem" data-checked="true">flour</li>
  //         <li data-type="taskItem" data-checked="true">baking powder</li>
  //         <li data-type="taskItem" data-checked="true">salt</li>
  //         <li data-type="taskItem" data-checked="false">sugar</li>
  //         <li data-type="taskItem" data-checked="false">milk</li>
  //         <li data-type="taskItem" data-checked="false">eggs
  //           <ul data-type="taskList">
  //             <li data-type="taskItem" data-checked="false">milk</li>
  //         </li>
  //         <li data-type="taskItem" data-checked="false">butter</li>
  //       <ul data-type="taskList">`
  // SetEditorContent(rich.value!,content)
}
onActivated(()=>{
  rich.value = RichText
  rich.value?.on('update', ({editor }) => {
    const content = JSON.stringify(editor.getJSON())
    result.value = content;
  })
})
onUnmounted(()=>{

})
const props = defineProps({
  callback: {
    type: Function,
    default: () => {}
  },
  data:{
    type: Object as PropType<Array<Task>>,
    
  }
})
const submit = () => {
  rich.value?.createNodeViews()
  props.callback(result.value)
}


</script>

const tes
<style  lang="less">
.tiptap {
  :first-child {
    margin-top: 0;
  }
  &:focus-visible{
    outline: none;
  }

  /* Heading styles */
  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    line-height: 1.1;
    margin-top: 2.5rem;
    text-wrap: pretty;
  }

  h1,
  h2 {
    margin-top: 3.5rem;
    margin-bottom: 1.5rem;
  }

  h1 {
    font-size: 1.4rem;
  }

  h2 {
    font-size: 1.2rem;
  }

  h3 {
    font-size: 1.1rem;
  }

  h4,
  h5,
  h6 {
    font-size: 1rem;
  }

  /* Code and preformatted text styles */
  code {
    background-color: var(--purple-light);
    border-radius: 0.4rem;
    color: var(--black);
    font-size: 0.85rem;
    padding: 0.25em 0.3em;
  }

  pre {
    background: var(--black);
    border-radius: 0.5rem;
    color: var(--white);
    font-family: 'JetBrainsMono', monospace;
    margin: 1.5rem 0;
    padding: 0.75rem 1rem;

    code {
      background: none;
      color: inherit;
      font-size: 0.8rem;
      padding: 0;
    }
  }

  mark {
    background-color: #FAF594;
    border-radius: 0.4rem;
    box-decoration-break: clone;
    padding: 0.1rem 0.3rem;
  }

  blockquote {
    border-left: 3px solid var(--gray-3);
    margin: 1.5rem 0;
    padding-left: 1rem;
  }

  hr {
    border: none;
    border-top: 1px solid var(--gray-2);
    margin: 2rem 0;
  }

  /* Task list specific styles */
  ul[data-type="taskList"] {
    list-style: none;
    margin-left: 0;
    padding: 0;

    li {
      align-items: flex-start;
      display: flex;

      > label {
        flex: 0 0 auto;
        margin-right: 0.5rem;
        user-select: none;
      }

      > div {
        flex: 1 1 auto;
      }
    }

    input[type="checkbox"] {
      cursor: pointer;
    }

    ul[data-type="taskList"] {
      margin: 0;
    }
  }
  .is-empty::before {
    color: gainsboro;
    content: attr(data-placeholder);
    float: left;
    height: 0;
    pointer-events: none;
  }
}
</style>