export interface questionInterface {
    id: string,
    QuestionUser: string,
    QuestionTitle: string,
    QuestionDescription: string
}

export interface createInterface {
    handleClose: () => void
    Open: boolean,
}
