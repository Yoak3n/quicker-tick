import {ExtractPropTypes} from 'vue'


export const calendarProps = {
    modelValue:{
        type: Date
    },
    yearAndMonth :{
        type: String ,
    }
}

export type CalendarProps = ExtractPropTypes<typeof calendarProps>

export const calendarEmits = {
    'update:modelValue': (value: Date) => value instanceof Date
}

export type CalendarEmits = typeof calendarEmits