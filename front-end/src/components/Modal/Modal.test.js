import { render } from '@testing-library/react';
import CustomModal from '../modal/Modal.component'


test('should render Modal component', () => {

    render(<CustomModal
        handleClose={test => void (test)}
        open={true}
    />);
})