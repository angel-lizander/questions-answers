import { DetailsInterface } from './Details.Interface'
import CustomModal from '../../components/Modal/Modal.component';
import { Accordion, AccordionSummary, Button, FormControl, Grid, TextField, Typography } from '@material-ui/core';
import styles from '../../components/Modal/Modal.style';
import { useEffect } from 'react';
import { GET, PUT } from './../../services/httpclient/httpclient.services'
import React, { useState } from 'react';
import SaveIcon from '@material-ui/icons/Save';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';




const Details = ({ handleClose, Open, Id }: DetailsInterface) => {

    const classes = styles();

    interface AnswerInterface {
        AnswerUser: string, AnswerDescription: string

    }

    var initialAnswer: AnswerInterface = {
        AnswerUser: "",
        AnswerDescription: ""
    };


    const [data, setData] = useState<any | null>(null);
    const [answer, setAnswer] = useState<AnswerInterface>(initialAnswer)

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setAnswer({ ...answer, [name]: value });
    }

    useEffect(() => {
        GET({ url: 'http://localhost:8080/questions/' + Id }).then(data => setData(data.data))
    }, [data, Id])


    const onAnswer = () => {

        var parameters = {
            QuestionAnswer: {
                AnswerUser: "",
                AnswerDescription: answer.AnswerDescription
            }
        };

        PUT({ url: 'http://localhost:8080/questions/' + Id, parameters: parameters }).then(response => response.success ? handleClose() : "")


    }

    return (


        <CustomModal open={Open} handleClose={() => handleClose()}>
            <Grid container className={classes.title}>
                <Grid item lg={5}>
                    Watch question
                </Grid>
            </Grid >
            <Grid container>
                <Grid item lg={8}>
                    <Grid item className={classes.items} lg={1} >
                        User
                    </Grid>
                    <Grid item lg={4}>
                        {data && data.Question.QuestionUser}
                    </Grid>
                </Grid>
            </Grid>
            <Grid container>
                <Grid item lg={8}>
                    <Grid item className={classes.items} lg={1} >
                        Title
                    </Grid>
                    <Grid item lg={4}>
                        {data && data.Question.QuestionTitle}
                    </Grid>
                </Grid>
            </Grid>
            <Grid container>
                <Grid item lg={8}>
                    <Grid item className={classes.items} lg={1} >
                        Description
                    </Grid>
                    <Grid item lg={4}>
                        {data && data.Question.QuestionDescription}
                    </Grid>
                </Grid>
            </Grid>
            <Accordion>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1a-content"
                    id="panel1a-header"
                >
                    <Typography className={classes.heading}>Answer</Typography>
                </AccordionSummary>
                <FormControl className={classes.formControl}>
                    <TextField
                        fullWidth
                        margin="dense"
                        multiline
                        rows="5"
                        variant="outlined"
                        id="additional-info"
                        defaultValue={data?.Question.QuestionAnswer.AnswerDescription}
                        name="AnswerDescription"
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleChange(e)}
                    />
                </FormControl>
                <Grid container xs={8}>
                    <Button
                        color="primary"
                        variant="contained"
                        endIcon={<SaveIcon />}
                        onClick={() => onAnswer()}
                    >
                        Answer
                    </Button>
                </Grid>

            </Accordion>

        </CustomModal >
    )

}


export default Details;