# gcr.io/deep-odds/tensorflow-serving
FROM tensorflow/serving:latest-devel

# Expose gRPC
EXPOSE 8500

RUN mkdir -p /models
COPY ./models /models

ENV MODEL_BASE_PATH=/models
ENV MODEL_NAME=model_67

RUN echo '#!/bin/bash \n\n tensorflow_model_server --port=8500 --model_name=${MODEL_NAME} \
--model_base_path=${MODEL_BASE_PATH}/${MODEL_NAME} "$@"' > /usr/bin/tf_serving_entrypoint.sh \
&& chmod +x /usr/bin/tf_serving_entrypoint.sh

ENTRYPOINT ["/usr/bin/tf_serving_entrypoint.sh"]