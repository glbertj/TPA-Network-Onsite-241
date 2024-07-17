from flask import Flask,request, jsonify
from flask_cors import CORS, cross_origin
from keras.models import load_model
import numpy as np
from collections import Counter
import librosa, librosa.display
import math
SAMPLE_RATE = 22050
N_MFCC = 13
N_FFT = 2048
HOP_LENGTH = 512
NUM_SEGMENTS = 10
num_samples_per_segment = int(SAMPLE_RATE * 5 / NUM_SEGMENTS)
expected = math.ceil(num_samples_per_segment / HOP_LENGTH)

app = Flask(__name__)
CORS(app,supports_credentials=True)   



@app.route('/getresult',methods=['POST'])
@cross_origin()
def members():
    if 'files' not in request.files:
        return jsonify({'error': 'No file part'}), 400
    file = request.files['files']
    if file.filename == '':
        return jsonify({'error': 'No selected file'}), 400
    if file :
        model = load_model('nice3.h5')
        result = predict_song(file,model)
        
        return jsonify({"result" : result}),200
    else:
        return jsonify({'error': 'File type not allowed'}), 400



def predict_song(file_path,model):
    signal, sr = librosa.load(file_path, sr=SAMPLE_RATE)
    num_segments = math.ceil(len(signal) / num_samples_per_segment)

    predictions = []

    for s in range(num_segments):
        start_sample = num_samples_per_segment * s
        end_sample = start_sample + num_samples_per_segment

        segment = signal[start_sample:end_sample]
        mfcc_features = librosa.feature.mfcc(y=segment, sr=SAMPLE_RATE, n_mfcc=N_MFCC, n_fft=N_FFT, hop_length=HOP_LENGTH)
        mfcc_features = mfcc_features.T[np.newaxis, ..., np.newaxis]

        prediction = model.predict(mfcc_features)
        predicted_label = np.argmax(prediction)
        predictions.append(predicted_label)

    label = ['ABBA_Dancing Queen', 'Bo Burnham_How the World Works', 'LOONA_Hi High',
 'Laufey_Promise', 'NewJeans_OMG', 'Phoebe Bridgers_I Know The End',
 'Phoebe Bridgers_That Funny Feeling', 'RADWIMPS_Zen Zen Zense',
 'Radiohead_No Surprises' ,'fourfolium_STEP by STEP UP']
    return label[Counter(predictions).most_common(1)[0][0]]

if __name__ == "__main__":
    app.run(debug=True)
    # app.run(host='0.0.0.0', port=5000)
    

    
    