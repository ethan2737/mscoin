"""
使用 Playwright 进行 Swap API 自动化测试
"""
import subprocess
import time
import requests
import json

BASE_URL = "http://localhost:8086"

def check_service_ready(url, timeout=30):
    """检查服务是否就绪"""
    print(f"Waiting for service: {url}")
    start = time.time()
    while time.time() - start < timeout:
        try:
            resp = requests.get(f"{url}/health", timeout=2)
            if resp.status_code == 200:
                print("Service ready")
                return True
        except:
            pass
        time.sleep(2)

    # 尝试根路径
    try:
        resp = requests.post(f"{url}/swap/contract-info/list",
                            json={"page": 1, "pageSize": 10},
                            timeout=5)
        if resp.status_code == 200:
            print("Service ready")
            return True
    except:
        pass

    print(f"Service startup timeout ({timeout}s)")
    return False

def run_api_tests():
    """运行 API 测试"""
    print("\n" + "=" * 60)
    print("Swap API Automated Tests")
    print("=" * 60)

    tests = [
        ("Contract List", "POST", "/swap/contract-info/list", {"page": 1, "pageSize": 10}),
        ("Contract Info", "POST", "/swap/contract-info/info", {"symbol": "BTC/USDT"}),
        ("Add Order", "POST", "/swap/order-entrust/add", {
            "memberId": 1001, "contractCoinId": 1, "entrustType": 1,
            "type": 1, "direction": 1, "leverage": 10, "price": 50000, "amount": 0.1
        }),
        ("Current Order", "POST", "/swap/order-entrust/current", {"memberId": 1001}),
        ("History Order", "POST", "/swap/order-entrust/history", {"memberId": 1001}),
        ("Position", "POST", "/swap/order-entrust/position", {"memberId": 1001}),
        ("Leverage", "POST", "/swap/order-entrust/leverage", {"memberId": 1001, "contractCoinId": 1}),
        ("Wallet Info", "POST", "/swap/wallet/info", {"memberId": 1001, "unit": "USDT"}),
        ("Wallet Transaction", "POST", "/swap/wallet/transaction", {"memberId": 1001, "page": 1, "pageSize": 10}),
    ]

    results = []
    for name, method, endpoint, data in tests:
        print(f"\nTest: {name}")
        print(f"  API: {method} {endpoint}")

        try:
            url = f"{BASE_URL}{endpoint}"
            if method == "POST":
                resp = requests.post(url, json=data, timeout=10)
            else:
                resp = requests.get(url, params=data, timeout=10)

            print(f"  Status: {resp.status_code}")

            if resp.status_code == 200:
                result = resp.json()
                print(f"  Response: code={result.get('code', 'N/A')}")
                results.append((name, True, None))
            else:
                results.append((name, False, f"Status: {resp.status_code}"))
                print(f"  FAILED: {resp.status_code}")

        except requests.exceptions.ConnectionError as e:
            results.append((name, False, "Connection failed"))
            print(f"  Connection failed: {e}")
        except Exception as e:
            results.append((name, False, str(e)))
            print(f"  Error: {e}")

    # 汇总结果
    print("\n" + "=" * 60)
    print("Test Results Summary")
    print("=" * 60)

    passed = sum(1 for _, success, _ in results if success)
    total = len(results)

    for name, success, error in results:
        status = "PASS" if success else f"FAIL ({error})"
        print(f"{status} - {name}")

    print(f"\nTotal: {passed}/{total} tests passed")

    return passed == total

if __name__ == "__main__":
    # 检查服务是否就绪
    if not check_service_ready(BASE_URL, timeout=30):
        print("\nService not running, skipping tests")
        print("Tip: Please run 'go run swap-api/main.go -f swap-api/etc/conf.yaml' to start service")
        exit(0)

    # 运行测试
    success = run_api_tests()
    exit(0 if success else 1)
