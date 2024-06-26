import { OrderCurrency } from "../../../../types/orderTypes";

interface OrderPriceProps {
  setOrderPrice: (value: number) => void;
  setOrderCurrency: (value: OrderCurrency) => void;
}

export default function OrderPrice(props: OrderPriceProps) {
  const handleCurrencyChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedCurrencyString = e.target.value; // Get the string value from the select element
    const selectedCurrency =
      selectedCurrencyString as keyof typeof OrderCurrency; // Type assertion to enum key
    const currencyValue = OrderCurrency[selectedCurrency]; // Get the enum value from the key

    if (currencyValue !== undefined) {
      // Ensure that the enum value is valid
      props.setOrderCurrency(currencyValue);
    }
  };

  return (
    <div>
      <label
        htmlFor="price"
        className="block font-medium leading-6 text-gray-900 after:content-['*'] after:ml-0.5 after:text-red-500"
      >
        Price
      </label>
      <div className="relative mt-2 rounded-md shadow-sm">
        <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
          <span className="text-gray-500 sm:text-sm">$</span>
        </div>
        <input
          onChange={(e) => props.setOrderPrice(Number(e.target.value))}
          type="text"
          name="price"
          id="price"
          className="block w-full rounded-md border-0 py-1.5 pl-7 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6 outline-indigo-300"
          placeholder="0.00"
        />
        <div className="absolute inset-y-0 right-0 flex items-center">
          <label htmlFor="currency" className="sr-only">
            Currency
          </label>
          <select
            id="currency"
            name="currency"
            className="h-full rounded-md border-0 bg-transparent py-0 pl-2 pr-7 text-gray-500 sm:text-sm sm:leading-6 outline-indigo-300"
            onChange={handleCurrencyChange}
          >
            <option value={OrderCurrency.UAH}>UAH</option>
            <option value={OrderCurrency.USD}>USD</option>
            <option value={OrderCurrency.EUR}>EUR</option>
          </select>
        </div>
      </div>
    </div>
  );
}
