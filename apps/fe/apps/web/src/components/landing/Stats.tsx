const stats = [
  { value: '500K+', label: 'Active Learners' },
  { value: '10K+', label: 'Courses Available' },
  { value: '98%', label: 'Completion Rate' },
  { value: '4.9/5', label: 'Average Rating' },
];

export function Stats() {
  return (
    <section className="py-16 bg-gray-50">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
          {stats.map((stat, index) => (
            <div key={index} className="text-center">
              <p className="text-4xl font-bold text-primary-600">{stat.value}</p>
              <p className="mt-2 text-sm text-gray-600">{stat.label}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}