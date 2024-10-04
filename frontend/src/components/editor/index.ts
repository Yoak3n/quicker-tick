import { Task } from '@/types'
import { model} from "../../../wailsjs/go/models";
function ParseTaskList(content:Task) {
    
}

import type{Editor} from '@tiptap/vue-3'
export function SetEditorContent(editor:Editor,content:any) {
    editor.commands.setContent(content)
}