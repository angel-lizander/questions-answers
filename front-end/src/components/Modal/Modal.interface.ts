export interface ModalInterface {
    open: boolean,    
    handleClose: () => void,
    children: React.ReactNode
}