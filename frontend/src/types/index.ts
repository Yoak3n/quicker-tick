interface ActionModel {
    id: string,
    name: string,
    description: string,
    icon: string,
    command: string,
    type: string,
}


interface TaskModel {
    id?: string,
    title?: string,
    description: string,
    status: number,
    priority: number,
    due_date: string,
    actions:Action[],
    tags: string[],
    children: TaskView[],
    root: boolean,
    checked: boolean,
}
export type TaskView = TaskModel
export type Action = ActionModel