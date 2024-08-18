import { Editor } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Text from '@tiptap/extension-text'

import TaskList from '@tiptap/extension-task-list'
import  TaskItem  from '@tiptap/extension-task-item'
import { Operation,taskDocument } from './extension'


export const  RichText  =  new Editor({
    content: "",
    extensions: [
      StarterKit.configure({
        document:false,
        text:false,
        bulletList:false
      }),
      taskDocument,
      TaskList,
      TaskItem.configure({
        nested:true,
      }),
      Operation,
      Text,
    ],
    autofocus:true,
    editable:true,
}) as Editor;