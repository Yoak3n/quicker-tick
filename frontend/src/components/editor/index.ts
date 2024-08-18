import { Task } from '@/types'

function ParseTaskList(content:Task) {
    
}

import type{Editor} from '@tiptap/vue-3'
export function SetEditorContent(editor:Editor,content:string) {
    editor.commands.setContent(content)
}