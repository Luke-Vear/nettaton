import { combineReducers } from 'redux'
import { NEW_QUESTION, ANSWER_QUESTION } from '../actions'

const question = (state = {}, action) => {
    switch (action.type) {
        case ANSWER_QUESTION:
            return state
        case NEW_QUESTION:
            return state
        default:
            return state
    }
}


const rootReducer = combineReducers({
    question
})

export default rootReducer