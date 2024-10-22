import { Task } from '@/types'
import { model} from "../../../wailsjs/go/models";
function ParseTaskList(content:Task) {
    
}

interface TaskModel {
    id?: string,
    title?: string,
    description: string,
    status: string,
    priority: string,
    dueDate: string,
    tags: string[],
    subtasks: TaskModel[],
    parentTask: TaskModel | null
}
export type TaskView = TaskModel
