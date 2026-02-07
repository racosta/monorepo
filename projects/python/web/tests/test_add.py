import re


def test_add(client):
    response = client.get("/")
    assert response.status_code == 200
    assert bool(re.match(r"[1-9]\d* \+ [1-9]\d* = [1-9]\d*", response.data.decode()))


if __name__ == "__main__":
    import pytest

    raise SystemExit(pytest.main([__file__]))
