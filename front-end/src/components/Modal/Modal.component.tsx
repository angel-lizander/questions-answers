import { Backdrop, Box, Fade, IconButton, Modal } from '@material-ui/core';
import { ModalInterface } from "./Modal.interface";

import styles from "./Modal.style";

import CloseIcon from '@material-ui/icons/Close';

const CustomModal = ({ handleClose, children, open }: ModalInterface) => {

  const classes = styles();

  return (
    <Modal
      open={open}
     
      aria-labelledby="transition-modal-title"
      aria-describedby="transition-modal-description"
      className={classes.modal}
      onClose={handleClose}
      closeAfterTransition
      disableEscapeKeyDown
      BackdropComponent={Backdrop}
      BackdropProps={{
        timeout: 500,
      }}
    >
      
      <Fade in={true}>
        <div className={classes.paper}>
          <Box display="flex" alignItems="center">
            <Box flexGrow={1} ></Box>
            <Box>
              <IconButton onClick={handleClose}>
                <CloseIcon />
              </IconButton>
            </Box>
          </Box>
          {children && children}
        </div>
      </Fade>
    </Modal>
  )
}

export default CustomModal;