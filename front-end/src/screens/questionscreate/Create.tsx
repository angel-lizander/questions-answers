import { createInterface, questionInterface } from './Create.interface'
import CustomModal from '../../components/Modal/Modal.component';
import { Button, FormControl, Grid, TextField } from '@material-ui/core';
import styles from '../../components/Modal/Modal.style';
import { POST } from './../../services/httpclient/httpclient.services'
import React, { useState } from 'react';
import SaveIcon from '@material-ui/icons/Save';


const CreateQuestion = ({ handleClose, Open }: createInterface) => {

    const classes = styles();
    var initialQuestion: questionInterface = {
        id: "",
        QuestionUser: "",
        QuestionTitle: "",
        QuestionDescription: ""
    };

    const [question, setQuestion] = useState<questionInterface>(initialQuestion);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setQuestion({ ...question, [name]: value });
    }

    const onCreate = () => {

        var parameters = {
            QuestionUser: question.QuestionUser,
            QuestionTitle: question.QuestionTitle,
            QuestionDescription: question.QuestionDescription
        };

        POST({ url: 'http://localhost:8080/questions/', parameters: parameters }).then(response => response.success ? handleClose() : "")


    }

    return (


        <CustomModal open={Open} handleClose={() => handleClose()}>
            <Grid container className={classes.title}>
                <Grid item lg={5}>
                    Create question
                </Grid>
            </Grid >
            <Grid container>
                <Grid item lg={4}>
                    <TextField
                        name="QuestionUser"
                        id="standard-basic"
                        label="User"
                        className={classes.textfield}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleChange(e)}
                    />
                </Grid>
            </Grid>
            <Grid item lg={6}>
                <TextField
                    name="QuestionTitle"
                    id="standard-basic"
                    label="Tittle"
                    className={classes.textfield}
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleChange(e)}
                />
            </Grid>
            <Grid item lg={12}>
                <Grid item lg={6}>
                    <FormControl className={classes.formControl}>
                        <TextField
                            fullWidth
                            margin="dense"
                            multiline
                            rows="5"
                            variant="outlined"
                            label="Question"
                            name="QuestionDescription"
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleChange(e)}
                            id="additional-info"
                        />
                    </FormControl>
                </Grid>
                <Button
                    color="primary"
                    variant="contained"
                    endIcon={<SaveIcon />}
                    onClick={() => onCreate()}
                >
                    Save
                </Button>
            </Grid>

        </CustomModal >
    )

}


export default CreateQuestion;