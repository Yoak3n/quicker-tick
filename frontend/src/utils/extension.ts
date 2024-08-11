import BulletList from '@tiptap/extension-bullet-list'
import TaskItem from '@tiptap/extension-task-item'

import Document from '@tiptap/extension-document'
export const taskDocument = Document.extend({
    content: 'taskList',
  })
export  const CustomTaskItem = TaskItem.extend({
    content: 'inline*',
    addKeyboardShortcuts(){
        return {
            "Tab": () => {
                console.log("enter tab");
                
                return this.editor.chain().focus().sinkListItem('taskItem').run()
            },
            "Shift+Tab": () => this.editor.chain().focus().liftListItem('taskItem').run(),
        }
    }
  })

export const Operation = BulletList.extend({
    addKeyboardShortcuts(){
        return {
            "Delete": () => this.editor.chain().focus().selectParentNode().deleteSelection().run(),
        }
    }
})