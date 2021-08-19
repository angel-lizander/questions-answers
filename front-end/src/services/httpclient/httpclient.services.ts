import { IHttpClient, IHttpResponse } from "./httpclient.interface";

export const GET = async ({ url }: IHttpClient) => {

    let ResponseHttp: IHttpResponse = { success: false, errorDetails: null, data: null };


    await fetch(url)
        .then(async response => {
            const data = await response.json();
            if (!response.ok) {

                const error = (data.data && data.data.message) || response.statusText;
                ResponseHttp.errorDetails = error;
                console.error(error);

            }
            ResponseHttp = { success: true, data: data, errorDetails: null }
        })
        .catch(error => {
            ResponseHttp.errorDetails = error;
            console.error('There was an error!', error);
        })

    return ResponseHttp
};

export const POST = async ({ url, parameters }: IHttpClient) => {

    let ResponseHttp: IHttpResponse = { success: false, errorDetails: null, data: null };


    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(parameters)
    };

    await fetch(url, requestOptions)
        .then(async response => {
            const isJson = response.headers.get('content-type')?.includes('application/json');
            const data = isJson && await response.json();
            if (!response.ok) {
                const error = (data && data.message) || response.status;
                ResponseHttp.errorDetails = error;
                console.error(error)
                return;
            }

            ResponseHttp = { success: true, data: data ? data : "", errorDetails: null }

        })
        .catch(error => {
            ResponseHttp.errorDetails = error;
            console.error('There was an error!', error);
        });


    return ResponseHttp;
}

export const PUT = async ({ url, parameters }: IHttpClient) => {


    const requestOptions = {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(parameters)
    };

    let ResponseHttp: IHttpResponse = { success: true, errorDetails: null, data: null };

    await fetch(url, requestOptions)
        .then(async response => {
            const isJson = response.headers.get('content-type')?.includes('application/json');

            if (!response.ok) {
                const data = isJson && await response.json();

                const error = (data && data.message) || response.status;
                console.error(error)
                ResponseHttp.errorDetails = error;
                ResponseHttp.success = false;

            }

        })
        .catch(error => {
            ResponseHttp.errorDetails = error;
            console.error('There was an error!', error);
        });

    return ResponseHttp;

}

export const DELETE = async ({ url }: IHttpClient) => {
    
    let ResponseHttp: IHttpResponse = { success: true, errorDetails: null, data: null };
    await fetch(url, { method: 'DELETE' })
        .then(async response => {

            if (!response.ok) {
                const data = await response.json();
                const error = (data && data.message) || response.status;
                ResponseHttp.errorDetails = error;
            }

        })
        .catch(error => {
            ResponseHttp.errorDetails = error;
            console.error('There was an error!', error);
        });

    return ResponseHttp;
};

