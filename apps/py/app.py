from flask import Flask, escape, request

import functools
import time
import os


def timer(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.perf_counter()
        value = func()
        end_time = time.perf_counter()
        run_time = end_time - start_time
        print(f"Finished {func.__name__!r} in {run_time:.4f} secs")
        return value

    return wrapper



if __name__ == "__main__":
    app = Flask(__name__)
    port = "8080"

    @app.route('/')
    @timer
    def hello():
        name = request.args.get("name","world")
        return f'Hello {escape(name)}'

    app.run(debug=True,host="0.0.0.0", port=port if os.environ.get("PORT") != "" else os.environ.get("PORT"))
